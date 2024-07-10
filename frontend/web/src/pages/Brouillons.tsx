import { DeleteIcon, ExternalLinkIcon } from '@chakra-ui/icons';
import {
	CircularProgress,
	Grid,
	GridItem,
	HStack,
	// Modal,
	// ModalBody,
	// ModalCloseButton,
	// ModalContent,
	// ModalHeader,
	// ModalOverlay,
	Tag,
	Tooltip,
	VStack,
} from '@chakra-ui/react';
import ArticleCard from 'components/Cards/ArticleCard';
// import Editor from 'components/Editor/Editor';
import SearchInput from 'components/Inputs/SearchInput';
import { useUIContext } from 'contexts/ui';
import { useUserContext } from 'contexts/user';
import * as React from 'react';
import { useEffect, useState } from 'react';

const Brouillons = (): JSX.Element => {
	const [search, setSearch] = useState('');
	// const [editor, setEditor] = useState<boolean>(false);
	// const [draft, setDraft] = useState({
	// 	title: '',
	// 	topic: '',
	// 	content: '',
	// });
	const { handleToast } = useUIContext();
	const { data, methods } = useUserContext();

	const uiLoadWrittenArticles = async () => {
		try {
			const res = await methods.articles.load.written();
			handleToast(res);
		} catch (error) {
			console.error(error);
		}
	};

	const uiDeleteArticle = async (articleId: number) => {
		try {
			const res = await methods.articles.delete({ id: articleId });
			handleToast(res, true);
		} catch (error) {
			console.error(error);
		}
	};

	const uiUpdateArticle = async (articleId: number) => {
		try {
			const res = await methods.articles.update({ id: articleId, newDraft: false });
			handleToast(res, true);
		} catch (error) {
			console.error(error);
		}
	};

	useEffect(() => {
		uiLoadWrittenArticles();
	}, []);

	if (!data.user.articles.written) {
		return (
			<>
				<VStack w="100%" h="100vh" justify="center">
					<CircularProgress size="120px" isIndeterminate color="black" />
				</VStack>
			</>
		);
	}

	return (
		<>
			<VStack w="100%" spacing={{ base: '8px', md: '12px', lg: '16px' }} align="start">
				<SearchInput
					value={search}
					inputId="brouillons-search-input"
					w={{ base: '100%', xl: '640px' }}
					placeholder="Cherchez parmis vos articles publiés"
					onChange={(e) => setSearch(e.target.value)}
					variant="primary-1"
				/>
				<HStack>
					<Tag bg="primary.yellow">
						{
							data.user.articles.written
								.filter((a) => a.draft)
								.filter((p) => (search !== '' ? p.title.includes(search) : true)).length
						}{' '}
						brouillon
						{data.user.articles.written.filter((a) => a.draft).length !== 1 && 's'}
					</Tag>
					<Tag bg="primary.blue">
						{data.user.articles.written
							.filter((a) => a.draft)
							.filter((p) => (search !== '' ? p.title.includes(search) : true))
							.map((p) => p.totalLikes)
							.reduce((a, v) => a + v, 0)}{' '}
						like
						{data.user.articles.written
							.filter((a) => a.draft)
							.filter((p) => (search !== '' ? p.title.includes(search) : true))
							.map((p) => p.totalLikes)
							.reduce((a, v) => a + v, 0) !== 1 && 's'}
					</Tag>
				</HStack>
				<Grid
					templateColumns={{ base: 'repeat(1, 1fr)', md: 'repeat(2, minmax(0, 1fr));' }}
					gap={{ base: 2, lg: 4 }}
					w="100%"
				>
					{data.user.articles.written
						.filter((a) => a.draft)
						.filter((p) => (search !== '' ? p.title.includes(search) : true))
						.map((brouillon, index) => (
							<GridItem key={`${index.toString()}`}>
								<ArticleCard
									id={brouillon.id}
									title={brouillon.title}
									// TODO: author name
									author="Author"
									date={new Date().toLocaleDateString('fr-FR')}
									// TODO: topic
									topic="Topic"
									content={brouillon.content}
									actions={[
										// <Tooltip label="Éditer l'article">
										// 	<span>
										// 		<EditIcon
										// 			onClick={() => {
										// 				setEditor(true);
										// 				setDraft({
										// 					title: brouillon.Title,
										// 					topic: brouillon.Topic,
										// 					content: brouillon.Content,
										// 				});
										// 			}}
										// 			color="black"
										// 		/>
										// 	</span>
										// </Tooltip>,
										<Tooltip label="Publier l'article">
											<span>
												<ExternalLinkIcon onClick={() => uiUpdateArticle(brouillon.id)} color="black" />
											</span>
										</Tooltip>,
										<Tooltip label="Supprimer définitivement">
											<span>
												<DeleteIcon onClick={() => uiDeleteArticle(brouillon.id)} color="black" />
											</span>
										</Tooltip>,
									]}
									view="publisher"
									views={brouillon.totalViews}
									likes={+brouillon.totalLikes}
								/>
							</GridItem>
						))}
				</Grid>
				{/* <Modal isOpen={editor} size="full" onClose={() => setEditor(false)}>
					<ModalOverlay />
					<ModalContent bg="black">
						<ModalHeader color="gray.100">Brouillon</ModalHeader>
						<ModalCloseButton />
						<ModalBody>
							<Editor
								placeholderTitle={draft.title}
								placeholderTopic={draft.topic}
								placeholderContent={draft.content}
							/>
						</ModalBody>
					</ModalContent>
				</Modal> */}
			</VStack>
		</>
	);
};

export default Brouillons;
